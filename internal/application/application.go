package application

import (
	"log"

	"github.com/sgash708/f-clock-io/internal/client/chromedriver"
	"github.com/sgash708/f-clock-io/internal/config"
	"golang.org/x/xerrors"
)

func RunClockIn() (err error) {
	d := chromedriver.StartDriver(config.Conf.Secret.URL)

	defer func() {
		var err error
		if err = d.Stop(); err != nil {
			return
		}
	}()

	site := diChromeDriver()
	if err := site.Login(); err != nil {
		log.Println(err)
		return err
	}
	if err := site.ClockIn(); err != nil {
		log.Println(err)
		return err
	}
	now := site.GetTime()
	if now == "" {
		log.Println("not exist time")
		return xerrors.New("not exist time.")
	}
	log.Printf("clock in: %v\n", now)

	return nil
}

func RunClockOut() (err error) {
	d := chromedriver.StartDriver(config.Conf.Secret.URL)

	defer func() {
		var err error
		if err = d.Stop(); err != nil {
			return
		}
	}()

	site := diChromeDriver()
	if err := site.Login(); err != nil {
		log.Println(err)
		return err
	}
	if err := site.NoteMemo(); err != nil {
		log.Println(err)
		return err
	}
	if err := site.ClockOut(); err != nil {
		log.Println(err)
		return err
	}
	now := site.GetTime()
	if now == "" {
		log.Println("not exist time")
		return xerrors.New("not exist time.")
	}
	log.Printf("clock out: %v\n", now)

	return nil
}

func RunStartBreak() (err error) {
	d := chromedriver.StartDriver(config.Conf.Secret.URL)

	defer func() {
		var err error
		if err = d.Stop(); err != nil {
			return
		}
	}()

	site := diChromeDriver()
	if err := site.Login(); err != nil {
		log.Println(err)
		return err
	}
	if err := site.StartBreak(); err != nil {
		log.Println(err)
		return err
	}
	now := site.GetTime()
	if now == "" {
		log.Println("not exist time")
		return xerrors.New("not exist time.")
	}
	log.Printf("break start: %v\n", now)

	return nil
}

func RunFinishBreak() (err error) {
	d := chromedriver.StartDriver(config.Conf.Secret.URL)

	defer func() {
		var err error
		if err = d.Stop(); err != nil {
			return
		}
	}()

	site := diChromeDriver()
	if err := site.Login(); err != nil {
		log.Println(err)
		return err
	}
	if err := site.FinishBreak(); err != nil {
		log.Println(err)
		return err
	}
	now := site.GetTime()
	if now == "" {
		log.Println("not exist time")
		return xerrors.New("not exist time.")
	}
	log.Printf("break finish: %v\n", now)

	return nil
}

func RunFixMemo() (err error) {
	d := chromedriver.StartDriver(config.Conf.Secret.URL)

	defer func() {
		var err error
		if err = d.Stop(); err != nil {
			return
		}
	}()

	site := diChromeDriver()
	if err := site.Login(); err != nil {
		log.Println(err)
		return err
	}
	if err := site.NoteMemo(); err != nil {
		log.Println(err)
		return err
	}
	now := site.GetTime()
	if now == "" {
		log.Println("not exist time")
		return xerrors.New("not exist time.")
	}
	log.Printf("fix note: %v\n", now)

	return nil
}

func diChromeDriver() chromedriver.ISite {
	return chromedriver.NewSite(
		config.Conf.Secret.Email,
		config.Conf.Secret.Password,
		config.Conf.Secret.UserID,
		config.Conf.Secret.Memo,
	)
}
