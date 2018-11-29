#include "ReadData.h"

SharedMemory getSharedMemory(string memName) {
  // Open the memory-mapped file
  HANDLE fileHandle = OpenFileMapping( PAGE_READONLY, FALSE, memName );
  if (fileHandle == NULL)
  {
    printf( "Could not open file mapping object (%d).\n", GetLastError() );
    return NULL;
  }

  // Get the data structure
  const SharedMemory* sharedData = (SharedMemory*)MapViewOfFile( fileHandle, PAGE_READONLY, 0, 0, sizeof(SharedMemory) );
  SharedMemory* localCopy = new SharedMemory;
  if (sharedData == NULL)
  {
    printf( "Could not map view of file (%d).\n", GetLastError() );

    CloseHandle( fileHandle );
    return localCopy;
  }

  // Ensure we're sync'd to the correct data version
  if ( sharedData->mVersion != SHARED_MEMORY_VERSION )
  {
    printf( "Data version mismatch\n");
    return localCopy;
  }

  unsigned int updateIndex(0);
  unsigned int indexChange(0);

  if ( sharedData->mSequenceNumber % 2 )
  {
    // Odd sequence number indicates, that write into the shared memory is just happening
    continue;
  }

  indexChange = sharedData->mSequenceNumber - updateIndex;
  updateIndex = sharedData->mSequenceNumber;

  //Copy the whole structure before processing it, otherwise the risk of the game writing into it during processing is too high.
  memcpy(localCopy,sharedData,sizeof(SharedMemory));


  if (localCopy->mSequenceNumber != updateIndex )
  {
    // More writes had happened during the read. Should be rare, but can happen.
    continue;
  }

  return localCopy;
}
