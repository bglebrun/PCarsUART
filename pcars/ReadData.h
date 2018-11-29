#ifndef _READ_DATA_HPP_
#define _READ_DATA_HPP_

// Used for memory-mapped functionality
#include <windows.h>
#include "sharedmemory.h"

// Used for this example
#include <stdio.h>
#include <conio.h>

SharedMemory getSharedMemory(string memName);

#endif
