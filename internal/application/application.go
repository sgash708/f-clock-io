package application

import (
	"log"

	"github.com/sgash708/f-clock-io/internal/client/chromedriver"
	"github.com/sgash708/f-clock-io/internal/config"
)

func RunClockIn() (err error) {
	d := chromedriver.StartDriver(config.Conf.Secret.URL)

	defer func() {
		var err error
		if err = d.Stop(); err != nil {
			return
		}
	}()

	site := chromedriver.NewSite(
		config.Conf.Secret.Email,
		config.Conf.Secret.Password,
		config.Conf.Secret.UserID,
	)
	if err := site.Login(); err != nil {
		log.Println(err)
		return err
	}

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

	site := chromedriver.NewSite(
		config.Conf.Secret.Email,
		config.Conf.Secret.Password,
		config.Conf.Secret.UserID,
	)
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

	site := chromedriver.NewSite(
		config.Conf.Secret.Email,
		config.Conf.Secret.Password,
		config.Conf.Secret.UserID,
	)
	if err := site.Login(); err != nil {
		log.Println(err)
		return err
	}

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

	site := chromedriver.NewSite(
		config.Conf.Secret.Email,
		config.Conf.Secret.Password,
		config.Conf.Secret.UserID,
	)
	if err := site.Login(); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
