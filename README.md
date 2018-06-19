What?
=====

(eventually) a tar-like tool for extracting stuff from FAT filesystems.

I'm working on this for my own personal learning and interest - please do not
use this as a source of truth for how to work with FAT, and please do not use
this in any kind of production environment.

Can't be bothered to put a LICENSE file in, so: All code is MIT licensed and
if I have copyright over the FAT data stored in the testdata folders etc. then
those are too. Please heed especially the part in the MIT license about there
being no warranty for this software.

Progress
========

FAT32
-----
* [x] Correctly parse LFN entries
* [x] Correctly parse a set of LFN entries
* [ ] Correctly parse file entries
* [ ] Correctly read 8.3 file entries from a directory
* [ ] Correctly read a directory
* [ ] Given a mock cluster map & fat header, correctly read a file
* [ ] Correctly read the root directory of an actual filesystem
* [ ] Correctly read a file from an actual filesystem

Tool
----
* [ ] list contents of a dir (`unfat -t -f blah.fat32 /path/to/dir`)
* [ ] extract one or more files (`unfat -x -f blah.fat32 /path/to/file/or/dir`)

Stretch goals
-------------

* [ ] Extend support to exFAT
* [ ] Extend support to FAT16
* [ ] Extend support to FAT12
* [ ] Add a file to a FAT filesystem
