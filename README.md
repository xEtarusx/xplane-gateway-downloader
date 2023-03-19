# X-Plane Gateway Downloader

Don't want to wait for the next X-Plane update for new version of airports?

With this console application you can download airport updates from the X-Plane Gateway (https://gateway.x-plane.com/) directly 
into your ``Custom Scenery`` folder.

## Table of contents

* [Setup](#setup)
* [How to use](#how-to-use)
  + [Install a new airport](#install-a-new-airport)
  + [Update installed airports](#update-installed-airports)
  + [Uninstall an airport](#uninstall-an-airport)


## Setup

### Step 1:

Download the latest version for your operating system:
* [Windows](https://github.com/xEtarusx/xplane-gateway-downloader/releases/latest/download/xplane-gateway-downloader-windows.zip) 
* [Mac OS](https://github.com/xEtarusx/xplane-gateway-downloader/releases/latest/download/xplane-gateway-downloader-mac_os.zip)
* [Linux 32-bit](https://github.com/xEtarusx/xplane-gateway-downloader/releases/latest/download/xplane-gateway-downloader-linux-32bit.zip) 
* [Linux 64-bit](https://github.com/xEtarusx/xplane-gateway-downloader/releases/latest/download/xplane-gateway-downloader-linux-64bit.zip) 


### Step 2:

Extract the downloaded .zip archive.

**Important**: Do **NOT** edit the ``config.json`` manually. Only use commands to interact with the application.

### Step 3:

Open a command line terminal in the folder where you extracted the application.

* Windows:
  * Shift + Right click -> "Open PowerShell window here"
* Mac OS:
  * Finder -> Services -> "New Terminal at Folder"
* Linux:
  * you should already know how to do this ;)

### Step 4:

Now you can run the application once to see if everything works correctly

* Windows:
  * ``.\xplane-gateway-downloader.exe``
* Mac OS & Linux:
  * ``./xplane-gateway-downloader``

If you followed correctly, you should see the application help page after running the command above.

### Step 5:

Set your current X-Plane version so the application knows which default airports you currently have.

**Note**: The commands are for the version 11.55 as an example. If you have a different version change 11.55 to your version.

* Windows:
  * ``.\xplane-gateway-downloader.exe config -v 11.55``
* Mac OS & Linux:
  * ``./xplane-gateway-downloader config -v 11.55``

### Step 6:

Set the path to the CustomScenery folder of X-Plane. All airports will be downloaded into this folder.

**Note**: Change the path to your "Custom Scenery" folder located in the X-Plane folder as needed.

* Windows:
  * ``.\xplane-gateway-downloader.exe config -csf "D:\path-to\X-Plane 12\Custom Scenery"``
* Mac OS & Linux:
  * ``./xplane-gateway-downloader config -csf "D:\path-to\X-Plane 12\Custom Scenery"``


## How to use

### Install a new airport

To install a new airport you need to execute the following command. Replace EDDF with any ICAO code you want to download.

* Windows:
  * ``.\xplane-gateway-downloader.exe install -i EDDF``
* Mac OS & Linux:
  * ``./xplane-gateway-downloader install -i EDDF``

### Update installed airports

This command updates all airports you previously installed with the command above.

* Windows:
  * ``.\xplane-gateway-downloader.exe update``
* Mac OS & Linux:
  * ``./xplane-gateway-downloader update``

### Uninstall an airport

If you would like to remove an airport which you installed before you can run the following command. 
Replace EDDF with any ICAO code of an airport you want to remove.

* Windows:
  * ``.\xplane-gateway-downloader.exe uninstall -i EDDF``
* Mac OS & Linux:
  * ``./xplane-gateway-downloader uninstall -i EDDF``
