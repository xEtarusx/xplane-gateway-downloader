# X-Plane Gateway Downloader

With this console application you can download airport updates from the X-Plane Gateway directly into your CustomScenery folder.

## Commands
```
$ xplane-gateway-downloader --help
NAME:
   X-Plane Gateway Downloader - Download airports from X-Plane Gateway with ease

USAGE:
   xplane-gateway-downloader [global options] command [command options] [arguments...]

DESCRIPTION:
   Update airport sceneries from the X-Plane Gateway

COMMANDS:
   install    Install a new airport scenery pack
   update     Update all installed airport scenery packs
   uninstall  Uninstall an installed airport scenery pack
   config     Configure the application
   help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```

### Configuration

Before you can manage new airports you have to set up the application.

```
$ xplane-gateway-downloader config --help
NAME:
   xplane-gateway-downloader config - Configure the application

USAGE:
   xplane-gateway-downloader config [command options] [arguments...]

OPTIONS:
   --custom-scenery-folder path, --csf path  The path to CustomScenery folder of x-plane
   --x-plane-version version, -v version     Set the current version of x-plane
```

Run the following commands and replace the placeholder.

- Set the path to the CustomScenery folder in the x-plane game folder. The path must contain the ending slashes!
  - **Windows:** (use double backslash due to character escaping)
    ```
    xplane-gateway-downloader config --custom-scenery-folder C:\\path\\to\\X-Plane 11\\Custom Scenery\\
    ```
  - **Linux:**
    ```
    xplane-gateway-downloader config --custom-scenery-folder /path/to/Custom Scenery/
    ```

- Set your X-Plane game version. This will allow the application to not download airports if there is no newer version available and therefore save disk space.
    Most of you will be at the newest version which currently is 11.55. If not, change the version in the command below.
    ```
    xplane-gateway-downloader config --x-plane-version 11.55
    ```

### Install an airport

```
$ xplane-gateway-downloader install --help
NAME:
   xplane-gateway-downloader install - Install a new airport scenery pack

USAGE:
   xplane-gateway-downloader install [command options] [arguments...]

OPTIONS:
   --icao ICAO, -i ICAO  Install an airport by ICAO code
```

Example: ``xplane-gateway-downloader install --icao EDDF``

### Update all installed airports

```
$ xplane-gateway-downloader update --help
NAME:
   xplane-gateway-downloader update - Update all installed airport scenery packs

USAGE:
   xplane-gateway-downloader update [command options] [arguments...]

OPTIONS:
   --help, -h  show help (default: false)
```

Example: ``xplane-gateway-downloader update``

### Uninstall an installed airport

```
$ xplane-gateway-downloader uninstall --help
NAME:
   xplane-gateway-downloader uninstall - Uninstall an installed airport scenery pack

USAGE:
   xplane-gateway-downloader uninstall [command options] [arguments...]

OPTIONS:
   --icao ICAO, -i ICAO  Uninstall an airport by ICAO code
```

Example: ``xplane-gateway-downloader uninstall --icao EDDF``