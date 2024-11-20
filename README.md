<p align="center">
    <img src="https://user-images.githubusercontent.com/8983173/130322857-185831e2-f041-46eb-a17f-0a69d066c4e5.png" alt="Gotenberg Logo" width="150" height="150" />
    <h3 align="center">Gotenberg</h3>
    <p align="center">A windows modification</p>
    <p align="center">
    I needed a version where you only need to generate PDF to HTML, so I stripped all the other functionality.
    This is an undockerized version of Gotenberg that works without Docker.
    </p>
</p>

---
## What is this?
A Windows single executable (exe) file to run Gotenberg in environments where no Docker is available. 
You can also build it for Linux; just set the `GOARCH` environment variable before building.

**BEWARE!!!**
Only the `/forms/chromium/convert/html/` endpoint works!

## Building
- Install Go 1.23
- Run `build.bat`

## Quick Start

Set `CHROMIUM_BIN_PATH` in your command line or as an environment variable.

```bash
set CHROMIUM_BIN_PATH=C:\Program Files\Google\Chrome\Application\chrome.exe
```

Run executable:

```bash
gotenberg.exe
```

I'm not fully clear on licensing and how it works, so please refer to the original repository for details.

