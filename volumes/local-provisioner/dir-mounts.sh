#!/bin/sh

mkdir /mnt/disks
for vol in vol1 vol2 vol3; do sudo mkdir /mnt/disks/$vol; sudo mount -t tmpfs $vol /mnt/disks/$vol; done
