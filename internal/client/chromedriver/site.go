package chromedriver

import "time"

type ISite interface {
	Login() error
	ClockIn() error
	ClockOut() error
	StartBreak() error
	FinishBreak() error
	NoteMemo() error
}

type Site struct {
	Email    string
	Password string
	UserID   string
}

const Memo = "[FE]API追加機能開発・リモート"

func NewSite(email, password, userID string) ISite {
	return &Site{Email: email, Password: password, UserID: userID}
}

func (s *Site) Login() error {
	// Fill
	if err := page.FindByXPath(`//*[@id="user_email"]`).Fill(s.Email); err != nil {
		return err
	}
	if err := page.FindByXPath(`//*[@id="login-page"]/main/div/form/div/div[3]/input[1]`).Fill(s.Password); err != nil {
		return err
	}
	// Click
	if err := page.FindByXPath(`//*[@id="login-page"]/main/div/form/div/div[5]/input`).Click(); err != nil {
		return err
	}

	return nil
}

func (s *Site) ClockIn() error {
	if err := page.FindByXPath(`//*[@id="contents"]/main/div/div[1]/div[1]/section/div/div[4]/button`).Click(); err != nil {
		return err
	}

	return nil
}

func (s *Site) ClockOut() error {
	if err := page.FindByXPath(`//*[@id="contents"]/main/div/div[1]/div[1]/section/div/div[4]/button[1]`).Click(); err != nil {
		return err
	}

	return nil
}

func (s *Site) StartBreak() error {
	return nil
}

func (s *Site) FinishBreak() error {
	return nil
}

func (s *Site) NoteMemo() error {
	time.Sleep(1 * time.Second)

	if err := page.FindByXPath(`//*[@id="contents"]/main/div/div[1]/div[1]/section/div/div[5]/button[1]`).Click(); err != nil {
		return err
	}

	time.Sleep(2 * time.Second)

	// Clear
	taPath := `//*[@id="vb-Dialog_2"]/div/div/div/div[2]/textarea`
	if err := page.FindByXPath(taPath).SendKeys("CLEAR"); err != nil {
		return err
	}
	// Fill
	if err := page.FindByXPath(taPath).Fill(Memo); err != nil {
		return err
	}
	// Click
	if err := page.FindByXPath(`//*[@id="vb-Dialog_2"]/div/div/div/div[3]/div/div/div/div[1]/button`).Click(); err != nil {
		return err
	}

	time.Sleep(2 * time.Second)

	return nil
}
