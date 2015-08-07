# Blade

Build XCode image catalogs for app icons, universal images, and more - automatically, elegantly, and quickly.



* You can use Blade in existing projects to generate all of your image catalogs with no extra work. It will automagically generate new icons for you based on what's inside existing image catalogs.
* You can also use templates of image catalogs to generate your own catalogs from (you can find some included in `/templates`), to generate new image catalogs from.


## How does it work?

Given a source image, with a high enough resolution, you can generate every needed image size for iPhone, Apple Watch, Universal and any other format Apple will come up with in the future. 

This can be done because Blade uses __the same XCode image catalog__ configuration file, as its own configuration source - no new concept introduced.


Supported workflows:

* __Prototyping__ ad-hoc, while prototyping projects
* __Development__ build with Build Steps, transforming all of your source image assets to image catalogs
* __CI__ in your CI servers, either on OSX or Linux (though Linux can't compile code in this case, you can still use it to do image processing)




## Quick start

Download one of the binaries in `releases`, suitable for your platform (You can run Blade on OSX, Linux and even Windows)

#### Create a new catalog

```
$ blade --source=iTunesArtwork@2x.png --template=templates/watch.json --out=out/watch --catalog
```

Here we want to create an app icon image catalog for Apple Watch. We're using the biggest icon image we have (Typically it is the iTunes Artwork icon, where you must upload a 1024x1024 image).

I am using my Apple Watch `Contents.json` template (for image catalog configuration), and I'm dropping all generated assets - images, Contents.json in the `out/watch` folder.
To explicitly generate an image catalog I'm specifying the `--catalog` flag.

#### Update an existing catalog

```
$ blade -s iTunesArtwork@2x.png -t existing.xcasset -o existing.xcasset
```

In this scenario, we're doing the same as before, but not generating a new catalog. Blade is reading an existing image catalog, and updates all images from the source image in-place.

We're also using shorthand flags notation.


#### Use a Bladefile

The best way to use Blade, is to set up a local `Bladefile` for your entire project. Within it, specify all of your resources. Blade will pick it up automatically.

```
$ blade --init
Wrote Bladefile.
```

Then use Blade (use --verbose if you want logs) within the same folder where your `Bladefile` lives:

```
$ blade --verbose
INFO[0000] Found a local Bladefile.
INFO[0000] Bladefile contains 8 blade defs.
```



## Hacking on Blade

Clone this project, then `rm -rf .git`.

```bash
$ git clone https://github.com/jondot/go-cli-starter
$ mv go-cli-starter my-project && cd my-project
$ rm -rf .git # you should be within my-project
$ make build
```

`make` should be your entry point.

* `make build` - build
* `make test` - test
* `make dist` - build and package binaries for multiple platforms

Note: You should edit your binary name in `Rakefile`.


# Contributing

Fork, implement, add tests, pull request, get my everlasting thanks and a respectable place here :).


# Copyright

Copyright (c) 2014 [Dotan Nahum](http://gplus.to/dotan) [@jondot](http://twitter.com/jondot). See MIT-LICENSE for further details.



