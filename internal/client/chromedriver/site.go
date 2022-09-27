package chromedriver

import (
	"log"
	"time"
)

type ISite interface {
	Login() error
	ClockIn() error
	ClockOut() error
	StartBreak() error
	FinishBreak() error
	NoteMemo() error
	GetTime() string
}

type Site struct {
	Email    string
	Password string
	UserID   string
	Memo     string
}

func NewSite(email, password, userID, memo string) ISite {
	return &Site{
		Email:    email,
		Password: password,
		UserID:   userID,
		Memo:     memo,
	}
}

func (s *Site) Login() error {
	time.Sleep(3 * time.Second)
	// Fill
	if err := page.FindByXPath(`/html/body/div[3]/div/main/div/form/div/div[2]/input`).Fill(s.Email); err != nil {
		return err
	}
	if err := page.FindByXPath(`//*[@id="login-page"]/main/div/form/div/div[3]/input[1]`).Fill(s.Password); err != nil {
		return err
	}
	// Click
	if err := page.FindByXPath(`//*[@id="login-page"]/main/div/form/div/div[5]/input`).Click(); err != nil {
		return err
	}
	time.Sleep(2 * time.Second)

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
	if err := page.FindByXPath(`//*[@id="contents"]/main/div/div[1]/div[1]/section/div/div[4]/button[2]`).Click(); err != nil {
		return err
	}
	time.Sleep(500 * time.Millisecond)

	return nil
}

func (s *Site) FinishBreak() error {
	if err := page.FindByXPath(`//*[@id="contents"]/main/div/div[1]/div[1]/section/div/div[4]/button`).Click(); err != nil {
		return err
	}
	time.Sleep(500 * time.Millisecond)

	return nil
}

func (s *Site) NoteMemo() error {
	time.Sleep(1 * time.Second)

	if err := page.FindByXPath(`//*[@id="contents"]/main/div/div[1]/div[1]/section/div/div[5]/button[1]`).Click(); err != nil {
		return err
	}

	time.Sleep(2 * time.Second)

	// Clear
	fp := page.FindByXPath(`//*[@id="vb-Dialog_2"]/div/div/div/div[2]/textarea`)
	str, err := fp.Text()
	if err != nil {
		return err
	}
	if str != "" {
		log.Println("exist text in note")
		if err := fp.Clear(); err != nil {
			return err
		}
		if err := fp.SendKeys(""); err != nil {
			return err
		}
	}
	var memo string
	if s.Memo == "" {
		memo = "[FE/BE]API追加機能開発・リモート"
	} else {
		memo = s.Memo
	}
	// Fill
	if err := fp.Fill(memo); err != nil {
		return err
	}
	// Click
	if err := page.FindByXPath(`//*[@id="vb-Dialog_2"]/div/div/div/div[3]/div/div/div/div[1]/button`).Click(); err != nil {
		return err
	}

	time.Sleep(2 * time.Second)

	return nil
}

func (s *Site) GetTime() string {
	content := GetHTML()
	now := content.Find(`#contents > main > div > div:nth-child(1) > div.vb-loading.vb-loading--block.vb-mb100 > section > div > div:nth-child(4)`).Text()

	return now
}
