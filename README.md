<a name="readme-top"></a>

<!-- PROJECT SHIELDS -->
[![Version][version-shield]][version-url]
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![Apache 2.0 License][license-shield]][license-url]



<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/softc24/evotor-resto-go">
    <img src="docs/logo.png" alt="Logo" height="80">
  </a>

  <h3 align="center">API-клиент приложения <a href="https://market.evotor.ru/store/apps/06341a0a-a2d4-4d7f-a24f-fcc26531efb1">"Мой Ресторан"</a></h3>

  <p align="center">
    Клиент для использования API приложения "Мой Ресторан" на языке Go.
    <br />
    <!-- <a href="https://github.com/softc24/evotor-resto-go"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/softc24/evotor-resto-go">View Demo</a>
    · -->
    <a href="https://github.com/softc24/evotor-resto-go/issues">Report Bug</a>
    ·
    <a href="https://github.com/softc24/evotor-resto-go/issues">Request Feature</a>
  </p>
</div>

<!-- TABLE OF CONTENTS -->
- [About The Project](#about-the-project)
  - [Built With](#built-with)
- [Getting Started](#getting-started)
  - [Installation](#installation)
- [Usage](#usage)
- [Roadmap](#roadmap)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)


<!-- ABOUT THE PROJECT -->
## About The Project

<!-- [![Product Name Screen Shot][product-screenshot]](https://example.com) -->

Библиотека предоставляет доступ к <a href="https://resto.evotor.tech/docs/">API приложения "Мой Ресторан"</a> для сторонних разработок на языке Go.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Built With

* [![Golang][Golang]][Golang-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

### Installation

Рекомендованный способ установки с помощью `go get`:

```
$ go get github.com/softc24/evotor-resto-go
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

```go
package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	evotorrestogo "github.com/softc24/evotor-resto-go"
)

func main() {
	token := os.Getenv("TOKEN")
	if token == "" {
		log.Fatalln("Token is empty")
	}

	ctx := context.TODO()

    // создаем клиент
	client := evotorrestogo.Client{
		BaseURL: evotorrestogo.DevURL,
		Token:   token,
	}

    // получаем список торговых точек
	stores, err := client.SelectStores(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%+v\n", stores[0])

	storeId := stores[0].UUID
    // получаем меню торговой точки
	menu, err := client.SelectMenu(ctx, storeId)
	if err != nil {
		log.Fatalln(err)
	}

	product := menu[0]
	log.Printf("%+v\n", product)

    // создаем заказ
	order := evotorrestogo.MakeOrder(strconv.FormatInt(time.Now().UnixMilli(), 32), "Комментарий", evotorrestogo.Contacts{
		Phone: "79990001234",
	}, []evotorrestogo.OrderPosition{
		evotorrestogo.MakeOrderPosition(product.UUID, product.Name, product.Price+100, product.Price+50, 1000),
	})

	order, err = client.CreateOrder(ctx, storeId, order)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%+v\n", order)

    // проверяем состояние заказа
	order, err = client.GetOrder(ctx, storeId, order.UUID)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%+v\n", order)
}


```

<!-- _For more examples, please refer to the [Documentation](https://example.com)_ -->

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ROADMAP -->
## Roadmap

See the [open issues](https://github.com/softc24/evotor-resto-go/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- LICENSE -->
## License

Distributed under the Apache-2.0 license. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

Project Link: [https://github.com/softc24/evotor-resto-go](https://github.com/softc24/evotor-resto-go)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ACKNOWLEDGMENTS -->
<!-- ## Acknowledgments

Use this space to list resources you find helpful and would like to give credit to. I've included a few of my favorites to kick things off!

* [Choose an Open Source License](https://choosealicense.com)
* [GitHub Emoji Cheat Sheet](https://www.webpagefx.com/tools/emoji-cheat-sheet)
* [Malven's Flexbox Cheatsheet](https://flexbox.malven.co/)
* [Malven's Grid Cheatsheet](https://grid.malven.co/)
* [Img Shields](https://shields.io)
* [GitHub Pages](https://pages.github.com)
* [Font Awesome](https://fontawesome.com)
* [React Icons](https://react-icons.github.io/react-icons/search)

<p align="right">(<a href="#readme-top">back to top</a>)</p> -->



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/softc24/evotor-resto-go.svg?style=for-the-badge
[contributors-url]: https://github.com/softc24/evotor-resto-go/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/softc24/evotor-resto-go.svg?style=for-the-badge
[forks-url]: https://github.com/softc24/evotor-resto-go/network/members
[stars-shield]: https://img.shields.io/github/stars/softc24/evotor-resto-go.svg?style=for-the-badge
[stars-url]: https://github.com/softc24/evotor-resto-go/stargazers
[issues-shield]: https://img.shields.io/github/issues/softc24/evotor-resto-go.svg?style=for-the-badge
[issues-url]: https://github.com/softc24/evotor-resto-go/issues
[license-shield]: https://img.shields.io/github/license/softc24/evotor-resto-go.svg?style=for-the-badge
[license-url]: https://github.com/softc24/evotor-resto-go/blob/master/LICENSE.txt
[version-shield]: https://img.shields.io/github/go-mod/go-version/softc24/evotor-resto-go.svg?style=for-the-badge
[version-url]: https://pkg.go.dev/github.com/softc24/evotor-resto-go
[product-screenshot]: docs/screenshot.png
[Golang]: https://img.shields.io/badge/Golang-000000?style=for-the-badge&logo=go&logoColor=white
[Golang-url]: https://go.dev/