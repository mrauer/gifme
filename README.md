<!-- PROJECT LOGO -->
<p align="center">
  <a href="https://github.com/mrauer/gifme">
    <img src="images/logo.png" alt="Logo">
  </a>

  <h3 align="center">GifMe</h3>

  <p align="center">
    Video to animated GIF converter.
    <br />
    <br />
    <a href="https://github.com/mrauer/gifme/issues">Report Bug</a>
    ·
    <a href="https://github.com/mrauer/gifme/issues">Request Feature</a>
  </p>
</p>

<!-- TABLE OF CONTENTS -->
<summary><h2 style="display: inline-block">Table of Contents</h2></summary>
  <ol>
    <li><a href="#install-for-linux">Install for Linux</a></li>
    <li><a href="#install-for-mac">Install for Mac</a></li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#license">License</a></li>
  </ol>

<!-- INSTALL FOR LINUX -->
## Install for Linux

The sofware require the two following dependencies:

```sh
apt install ffmpeg
apt install gifsicle
```

Then here is how you can install `gifme`:

```sh
curl -Lo gitdump https://github.com/mrauer/gifme/releases/download/v1.0.0/gifme_1.0.0_linux_amd64 && chmod +x gifme && sudo mv gifme /usr/local/bin

```

<!-- INSTALL FOR MAC -->
## Install for Mac

Install the dependencies (we assume you have brew):

```sh
brew install ffmpeg
brew install gifsicle
```

Run the following command:

```sh
curl -Lo gitdump https://github.com/mrauer/gitdump/releases/download/v0.2.0/gifme_1.0.0_darwin_amd64 && chmod +x gifme && sudo mv gifme /usr/local/bin
```

<!-- USAGE -->
## Usage

Usage is platform agnostic:

```sh
gifme -input <path_to_your_mp4>
```

This will generate an optimized gif file named `output.gif` at the location you ran that command.

<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.
