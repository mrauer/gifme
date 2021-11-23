# GifMe

<!-- PROJECT LOGO -->
<p align="center">
  <a href="https://github.com/mrauer/gifme">
    <img src="images/logo.png" alt="Logo">
  </a>

  <p align="center">
    Converts your video to a GIF.
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
base_url='https://github.com/mrauer/gifme/releases/download/v1.0.0'
curl -Lo gifme "$base_url"/gifme_1.0.0_linux_amd64 && \
chmod +x gifme && sudo mv gifme /usr/local/bin
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
base_url='https://github.com/mrauer/gifme/releases/download/v1.0.0'
curl -Lo gifme "$base_url"/gifme_1.0.0_darwin_amd64 && \
chmod +x gifme && sudo mv gifme /usr/local/bin
```

<!-- USAGE -->
## Usage

Usage is platform agnostic:

```sh
gifme --input <path-to-your-video-file>
```

This will generate an optimized gif file named `output.gif` at the location you ran that command.

Currently only supporting `.mp4` files.

<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.
