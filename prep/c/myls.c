#include <stdio.h>
#include <stdbool.h>
#include <unistd.h>
#include <stdlib.h>
#include <errno.h>
#include <dirent.h>
#include <string.h>

#define MAX_PATH_SIZE 500
#define B 1
#define KB B*1024L
#define MB KB*1024L
#define GB MB*1024L
#define TB GB*1024L
#define PB TB*1024L


char *join_path(char *head, char *tail);
char *alloc_path();

typedef struct LsOpts {
    bool long_fmt;
    bool human;
} LsOpts;

int long_fmt(LsOpts opts, unsigned int num_paths, char **paths);
int default_fmt(LsOpts opts, unsigned int num_paths, char **paths);

int main(int argc, char *argv[]) {
    char *path = alloc_path();

    LsOpts *opts = malloc(sizeof(LsOpts));
    int ch;
    while ((ch = getopt(argc, argv, "lh")) != -1) {
        switch (ch) {
        case 'l':
            opts->long_fmt = true;
            break;
        case 'h':
            opts->human = true;
            break;
        case '?':
        default:
            continue;
        }
    }

    int num_paths = argc > optind ? argc - optind : 1;
    char **paths = malloc(num_paths);
    for (int i = 0; i < num_paths; i++) {
        *(paths+i) = alloc_path();
    }

    if (optind >= argc) {
        // Only path is .
        getcwd(*paths, MAX_PATH_SIZE);
    } else {
        for (int i = 0; i < num_paths; i++) {
            *(paths+i) = *(argv+optind+i);
        }
    }

    int status = 0;
    if (opts->long_fmt) {
        status = long_fmt(*opts, num_paths, paths);
    } else {
        status = default_fmt(*opts, num_paths, paths);
    }

    if (status) {
        exit(EXIT_FAILURE);
    } else {
        exit(EXIT_SUCCESS);
    }
}

int default_fmt(LsOpts opts, unsigned int num_paths, char **paths) {
    /*
     * Returns 0 on successful exit, non-zero on error.
     */
    for (int i = 0; i < num_paths; i++) {
        DIR *dp;
        struct dirent *dirp;
        char *path = *(paths+i);

        dp = opendir(path);
        if (dp == NULL) {
            fprintf(stderr, "%s: No such file or directory\n", path);
            return -1;
        }

        printf("%s:\n", path);

        while ((dirp = readdir(dp)) != NULL) {
            printf("%s\n", dirp->d_name);
        }

        closedir(dp);
    }

    return 0;
}

char *format_f_size(long size, bool human) {
  char *size_str;
  if (human) {
    if (size < KB) {
      asprintf(&size_str, "%ldB", size);
    } else if (size < MB) {
      asprintf(&size_str, "%.1fK", size / ((double) KB));
    } else if (size < GB) {
      asprintf(&size_str, "%.1fM", size / ((double) MB));
    } else if (size < TB) {
      asprintf(&size_str, "%.1fG", size / ((double) GB));
    } else if (size < PB) {
      asprintf(&size_str, "%.1fT", size / ((double) TB));
    } else {
      asprintf(&size_str, "%.1fP", size / ((double) PB));
    }
  } else {
    asprintf(&size_str, "%ld", size);
  }
  return size_str;
}

int long_fmt(LsOpts opts, unsigned int num_paths, char **paths) {
    /*
     * Returns 0 on successful exit, non-zero on error.
     */
    for (int i = 0; i < num_paths; i++) {
        DIR *dp;
        struct dirent *dirp;
        char *path = *(paths+i);

        dp = opendir(path);
        if (dp == NULL) {
            fprintf(stderr, "%s: No such file or directory\n", path);
            return -1;
        }

        printf("%s:\n", path);

        printf("%20s\t%s\n", "Name", "Size");
        printf("%20s\t%s\n", "----", "----");
        while ((dirp = readdir(dp)) != NULL) {
            char *full_path = join_path(path, dirp->d_name);
            FILE *fp = fopen(full_path, "r");
            fseek(fp, 0, SEEK_END);
            char *f_size = format_f_size(ftell(fp), opts.human);
            fclose(fp);
            printf("%20s\t%s\n", dirp->d_name, f_size);
        }

        closedir(dp);
    }

    return 0;
}

char *join_path(char *head, char *tail) {
    char *full_path = alloc_path();
    int i;
    for (i = 0; i < strlen(head); i++) {
        *(full_path+i) = *(head+i);
    }
    full_path[i++] = '/';
    for (int j = 0; j < strlen(tail); j++) {
        full_path[i++] = tail[j];
    }
    return full_path;
}

char *alloc_path() {
    return malloc(MAX_PATH_SIZE * sizeof(char));
}

/*
*
* Some ideas:
* [x] handle .
* [x] handle ..
* [x] error if the path is not found
* [x] print file size
* [x] support some basic flags, e.g. -lh
* [ ] hide "hidden" files by default, require -a argument
* [x] allow passing multiple arguments
* [ ] print columns that depend on the width of the terminal, and the width of the filenames.
* [ ] error for a path that's too long.
* [x] move to bradfield repo
* [ ] handle broken symlinks
* [ ] print usage when wrong args are passed
*/
