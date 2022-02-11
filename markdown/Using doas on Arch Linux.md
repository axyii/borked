[//]: # (2021-12-16)

# Using doas on Arch Linux

I recently uninstalled 'sudo' and started using 'doas': a utility used to assume the identity of another user. It's just like sudo but easier to manage and very simple to use. I would not recommend using it in systems where highly refined user permissions are required, but if you have a single user box or a desktop, I would say you are good to go. One of the major reasons I switched to is because it's wayy smaller and minimal than sudo.So, let's get started on setting it up


Install doas from the Arch repository as root

```
pacman -S doas
```

Create a simple config for doas in /etc/doas.conf as root

```
/etc/doas.conf
---------------

permit :wheel
```
or you could even use the persist feature which does not ask for a password after authentication.

Note: The persist feature is disabled by default because it is new and potentially dangerous. In the original doas by OpenBSD, a kernel API is used to set and clear timeouts. This API is OpenBSD specific and no similar API is available on other operating systems. As a workaround, the persist feature is implemented using timestamp files similar to sudo.

```
/etc/doas.conf
---------------

permit persist :wheel
```

Now, check the configuration for any errors using

```
doas -C /etc/doas.conf
```

To rectify any errors or for more configuration options read the man pages

```
man doas.conf
```

If you have sudo installed you can just uninstall it using

```
doas pacman -Rs sudo
```

Sometimes, makepkg in Arch Linux depends on sudo so you have to create a symlink to sudo which enables you to use the command sudo to invoke doas.

```
doas ln -s /usr/bin/doas /usr/bin/sudo
```

Now you have doas fully configured on your system.
