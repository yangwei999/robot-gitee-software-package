package pullrequestimpl

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"os/exec"
	"strings"

	"github.com/opensourceways/robot-gitee-software-package/softwarepkg/domain"
)

func (impl *pullRequestImpl) createPr() error {
	sigInfoFile := fmt.Sprintf("sig/%s/sig-info.yaml",
		impl.pkg.Application.ImportingPkgSig)
	sigInfoData, err := impl.genAppendSigInfoData()
	if err != nil {
		return err
	}

	subDirName := strings.ToLower(impl.pkg.Name[:1])
	newRepoFile := fmt.Sprintf("sig/%s/src-openeuler/%s/%s.yaml",
		impl.pkg.Application.ImportingPkgSig, subDirName, impl.pkg.Name)
	newRepoData, err := impl.genNewRepoData()
	if err != nil {
		return err
	}

	cmd := exec.Command(impl.cfg.ShellScript, impl.cfg.Robot.Username,
		impl.cfg.Robot.Password, impl.cfg.Robot.Email, impl.branchName(),
		impl.cfg.PR.Org, impl.cfg.PR.Repo,
		sigInfoFile, sigInfoData, newRepoFile, newRepoData,
	)

	if output, err := cmd.CombinedOutput(); err != nil {
		return errors.New(string(output))
	}

	return nil
}

func (impl *pullRequestImpl) branchName() string {
	return fmt.Sprintf("software_package_%s", impl.pkg.Name)
}

func (impl *pullRequestImpl) prName() string {
	return impl.pkg.Name + impl.cfg.PR.PRName
}

func (impl *pullRequestImpl) genAppendSigInfoData() (string, error) {
	data := struct {
		PkgName       string
		ImporterEmail string
		Importer      string
	}{
		PkgName:       impl.pkg.Name,
		ImporterEmail: impl.pkg.ImporterEmail,
		Importer:      impl.pkg.ImporterName,
	}

	return impl.genTemplate(impl.cfg.Template.AppendSigInfo, data)
}

func (impl *pullRequestImpl) genNewRepoData() (string, error) {
	data := struct {
		PkgName       string
		PkgDesc       string
		SourceCodeUrl string
		BranchName    string
		ProtectType   string
		PublicType    string
	}{
		PkgName:       impl.pkg.Name,
		PkgDesc:       impl.pkg.Application.PackageDesc,
		SourceCodeUrl: impl.pkg.Application.SourceCode.Address,
		BranchName:    impl.cfg.PR.NewRepoBranch.Name,
		ProtectType:   impl.cfg.PR.NewRepoBranch.ProtectType,
		PublicType:    impl.cfg.PR.NewRepoBranch.PublicType,
	}

	return impl.genTemplate(impl.cfg.Template.NewRepoFile, data)
}

func (impl *pullRequestImpl) genTemplate(fileName string, data interface{}) (string, error) {
	tmpl, err := template.ParseFiles(fileName)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	if err = tmpl.Execute(buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (impl *pullRequestImpl) getRobotLogin() (string, error) {
	if impl.robotLogin == "" {
		v, err := impl.cli.GetBot()
		if err != nil {
			return "", err
		}

		impl.robotLogin = v.Login
	}

	return impl.robotLogin, nil
}

func (impl *pullRequestImpl) submit() (dpr domain.PullRequest, err error) {
	robotName, err := impl.getRobotLogin()
	if err != nil {
		return
	}

	head := fmt.Sprintf("%s:%s", robotName, impl.branchName())
	pr, err := impl.cli.CreatePullRequest(
		impl.cfg.PR.Org, impl.cfg.PR.Repo, impl.prName(),
		impl.pkg.Application.ReasonToImportPkg, head, "master", true,
	)
	if err != nil {
		return
	}

	dpr = domain.PullRequest{
		Num:  int(pr.Number),
		Link: pr.HtmlUrl,
		Pkg:  impl.pkg.SoftwarePkgBasic,
	}

	return
}
