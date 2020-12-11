# Device Management

Here you can manage particular device.

## First look

As you click on any device you'll get to this page.

![Tabs](Images/device-registry/base.jpg?raw=true)

**Mark 1** - Refresh device data button.

**Mark 2** - Bulb color shows if device is enabled(green) or disabled(red), acts same at [Device Registry Page](Devices-Registry-Management-Page.md)

**Mark 3** - Device name

**Mark 4** - Device ID

**Mark 5** - Enter Edit Mode

## Edit Mode

After clicking on **Edit** button(**Mark 5**), you'll be able to edit device name and tags:

![Tabs](Images/device-registry/edit-mode.jpg?raw=true)


## State Card

After scrolling little bit down, you can see the Device State Card. It has two columns: **Reported** and **Desired** state:

![Tabs](Images/device-registry/state-base.jpg?raw=true)

**Mark 1 - Reported** state is the state received from the device.
Here you can see a last report timestamp and "version" - order number(**Marks 3 and 2**)

Same for **Desired**(**Mark 4**) and **Marks 5 and 6** for desired state version and timestamp.

By clicking on **Edit** button(**Mark 7**) - you enter **Desired** state edit mode(JSON editor - **Mark 1** below) - this is the data to be sent to the device.

![Tabs](Images/device-registry/state-edit-mode.jpg?raw=true)