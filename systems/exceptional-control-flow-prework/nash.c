#include <stdio.h>
#include <stdio.h>
#include <stdbool.h>
#include <unistd.h>
#include <stdlib.h>
#include <errno.h>
#include <dirent.h>
#include <string.h>
#include <time.h>


const unsigned short NUM_GHOST_QUOTES = 7;
const char* GHOST_QUOTES[7] = {
  "In one aspect, yes, I believe in ghosts, but we create them. We haunt ourselves.",
  "The house smelled musty and damp, and a little sweet, as if it were haunted by the ghosts of long-dead cookies.",
  "Now I know what a ghost is. Unfinished business, that's what.",
  "The people you love become ghosts inside of you, and like this you keep them alive.",
  "At night, here in the library, the ghosts have voices.",
  "Ghosts have a way of misleading you; they can make your thoughts as heavy as branches after a storm.",
  "The ghosts of things that never happened are worse than the ghosts of things that did.",
};

typedef struct NashOpts {
    char* cmd;
} NashOpts;

int main(int argc, char** argv) {
  NashOpts *opts;
  int ch;
  while ((ch = getopt(argc, argv, "c")) != -1) {
    switch (ch) {
      case 'c':
        // TODO: figure out how to properly initialize this struct so it doesn't segfault
        opts->cmd = optarg;
        /* printf("Received command \"%s\"", opts->cmd); */
        break;
      case '?':
      default:
        continue;
    }
  }

  char* prompt = "👻>>";
  int MAX_INPUT_SIZE = 500;
  char input[MAX_INPUT_SIZE];
  char* scan_result;
  // Use time as a seed to select a ghost quote for exit.
  time_t rawtime;
  time(&rawtime);
  srand(rawtime);
  do {
    printf("%s ", prompt);
    scan_result = fgets(input, MAX_INPUT_SIZE, stdin);
    printf("%s", input);
  } while (scan_result != NULL);
  printf("\n**%s**\n", GHOST_QUOTES[rand() % NUM_GHOST_QUOTES]);
  exit(EXIT_SUCCESS);
}
