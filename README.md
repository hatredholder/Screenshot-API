<div align="center">

# Screenshot-API

<div align="center">
    	<a href="#wrench-installation">Installation</a>
  <span> • </span>
       	<a href="#gear-usage">Usage</a>
  <span> • </span>
  	<a href="#stars-contributing">Contributing</a>
  <p></p>
</div> 

</div>

<div align="center">

**Screenshot-API** is a small [Docker](https://www.docker.com/) image with [chromedp](https://github.com/chromedp/chromedp) installed and **Golang**-based Webserver to interact with it. It can be used to take screenshots of provided websites.

</div>

# :wrench: Installation

Before running anything you need to make sure you have [Docker](https://www.docker.com/) installed.

1. Build a **Docker** container with:
```
make build
```

2. Run the **Docker** container with:
```
make run
```

# :gear: Usage
**Screenshot-API** will start on port 8080. To send a request to screenshot a website, for example, ```google.com```, go to:

```
http://localhost:8080/screenshot/?url=https://google.com
```

This will return a **JSON** response with a link to the screenshot image, in this format:
```
{"screenshotUrl":"localhost:8080/storage/LfUqsp.png"}
```

# :stars: Contributing

Contributions are welcome! Having an **issue**? An **idea** to improve existing functionality? **Critique**? 

Please leave **pull requests/issues** with your suggestions!
