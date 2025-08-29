#include <stdio.h>
#include <stddef.h>
#include <stdlib.h>
#include <stdarg.h>
#include <string.h>
#include <errno.h>
#include <title.h>
#include <limits.h>
#include <unistd.h>
#include <signal.h>
#include <sys/wait.h>
# include <errno.h>
# include <unistd.h>


extern char** environ;
void utTitle::my_initproctitle(char* argv[], char** last)
{
        int i = 0;
        char* p_tmp = NULL;
        size_t i_size = 0;
        //environ = last;
        for(i = 0; environ[i]; i++){
                i_size += strlen(environ[i]) + 1;
        }
 
        p_tmp = (char *)malloc(i_size);
        if(p_tmp == NULL){
                return ;
        }
 
        *last = argv[0];
        for(i = 0; argv[i]; i++){
                *last += strlen(argv[i]) + 1;
        }
 
        for(i = 0; environ[i]; i++){
                i_size = strlen(environ[i]) + 1;
                *last += i_size;
                strncpy(p_tmp, environ[i], i_size);
                environ[i] = p_tmp;
                p_tmp += i_size;
        }
 
        (*last)--;
 
        return ;
}
 
 
void utTitle::my_setproctitle(char* argv[], char** last, char* title)
{
        char* p_tmp = NULL;
        /* argv[1] = NULL; */
        p_tmp = argv[0];
        strncpy(p_tmp, title, *last - p_tmp);
        return ;
}

int utTitle::process_keepalive(void)
{
    while (true) {
        int pid = fork();
        if (pid < 0) {
            return -1;
        } else if (pid == 0 ) {
            return 0;
        } else {
            int status = 0;
            int ret = waitpid(pid, &status, 0);
            if (ret < 0) {
                if (1) {
                    exit(EXIT_SUCCESS);
                } else {
                    exit(EXIT_FAILURE);
                }
            }
            if (WIFEXITED(status)) {
                exit(EXIT_SUCCESS);
            } else if (WIFSIGNALED(status)) {
                usleep(1000 * 1000);
                continue;
            } else {
                exit(EXIT_FAILURE);
            }
        }
    }
    return -1;
}


