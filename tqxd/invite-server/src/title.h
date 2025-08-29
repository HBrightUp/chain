#ifndef _UT_TITLE_H_
#define _UT_TITLE_H_
class utTitle
{
public:
    utTitle() {}
    ~utTitle() {}
    /* should call this function in the beginning of main */
    static void my_initproctitle(char* argv[], char** last);

    /* update process titile */
    static void my_setproctitle(char* argv[], char** last, char* title);

    static int process_keepalive(void);
};

#endif
