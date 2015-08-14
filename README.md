# Blade ![Build Status](https://travis-ci.org/jondot/blade.svg?branch=master)

Automatically build and rebuild XCode image catalogs for app icons, universal images, and more.



* Use in existing projects to generate image catalogs with no extra work. Blade will automagically refresh your image catalogs based on given master images.
* Use templates of image catalogs to generate new catalogs (see [templates](templates/)).

![](docs/blade.gif)



## Quick start

Download one of the binaries in [releases](https://github.com/jondot/blade/releases), and put in your `PATH` or just include in each XCode project's root.


#### Use a Bladefile

The best way to use Blade, is to set up a local `Bladefile` for your entire project. Within it, specify all of your resources. Blade will pick it up automatically.

```
$ blade --init
Wrote Bladefile.
```
Here is how your `Bladefile` would look like:

```yaml
blades:
  - source: iTunesArtwork@2x.png
    mount: foobar/Assets.xcassets/AppIcon.appiconset
  - source: Spaceship_1024.png
    mount: foobar/Assets.xcassets/Spaceship.imageset
```

It was made for this project structure:

```
foobar
├── foobar
│   ├── AppDelegate.swift
│   ├── Assets.xcassets
│   │   ├── AppIcon.appiconset
│   │   │   └── Contents.json
│   │   └── Spaceship.imageset
│   │       ├── Contents.json
```

Then use Blade (use --verbose if you want logs) within the same folder where your `Bladefile` lives:

```
$ blade --verbose
INFO[0000] Found a local Bladefile.
INFO[0000] Bladefile contains 2 blade defs.
```

And it will generate all of the images needed within each image catalog.

To make this happen before each build see [how to run a script while building a product](https://developer.apple.com/library/ios/recipes/xcode_help-project_editor/Articles/AddingaRunScriptBuildPhase.html)




#### Use directly

```
$ blade --source=iTunesArtwork@2x.png --template=templates/watch.json --out=out/watch --catalog
```

Here we want to create an app icon image catalog for Apple Watch. We're using the biggest icon image we have (Typically it is the iTunes Artwork icon, where you must upload a 1024x1024 image).

I am using my Apple Watch `Contents.json` template (for image catalog configuration), and I'm dropping all generated assets - images, Contents.json in the `out/watch` folder.
To explicitly generate an image catalog I'm specifying the `--catalog` flag.


```
$ blade -s iTunesArtwork@2x.png -t existing.imageset -o existing.imageset
```

In this scenario, we're doing the same as before, but not generating a new catalog. Blade is reading an existing image catalog, and updates all images from the source image in-place.

We're also using shorthand flags notation.



## How does it work?

Given a source image, with a high enough resolution, you can generate every needed image size for iPhone, Apple Watch, Universal and any other format Apple will come up with in the future. 

This can be done because Blade uses __the same XCode image catalog__ configuration file, as its own configuration source - no new concept introduced.


Supported workflows:

* __Prototyping__ ad-hoc, while prototyping projects
* __Development__ build with Build Steps, transforming all of your source image assets to image catalogs
* __CI__ in your CI servers, either on OSX or Linux (though Linux can't compile code in this case, you can still use it to do image processing)




## Hacking on Blade


```bash
```

`make` should be your entry point.

* `make setup` - setup the tooling for the project
* `make` - default build
* `make release` - cross build a release for multiple platforms


# Contributing

Fork, implement, add tests, pull request, get my everlasting thanks and a respectable place here :).


# Copyright

Copyright (c) 2014 [Dotan Nahum](http://gplus.to/dotan) [@jondot](http://twitter.com/jondot). See MIT-LICENSE for further details.



