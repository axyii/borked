[//]: # (2022-06-27)

# Setting up pipewire in Arch Linux

Pipewire is a server for handling audio and video streams. It's mostly used for audio, so that is what I will be showing today.

First install pipewire, pipewire-pulse and pipewire-media-session

```
pacman -S pipewire pipewire-pulse pipewire-media-session
```
If you want a graphical manager for pipewire then install pavucontrol, which is usually recommended
```
pacman -S pavucontrol
```
Now restart your computer for the changes to take place.
