#include <curses.h>
#include <stdio.h>
#include "test.h"

int main() {
    InitKeyboard();

    printf("\nEnter: ");
    refresh();

    for (;;) {
        int r = GetCharacter();
        printf("%c", r);
        refresh();

        if (r == 'q') {
            break;
        }
    }

    CloseKeyboard();

    return 0;
}

void InitKeyboard() {
    initscr();
    noecho();
    cbreak();
    keypad(stdscr, TRUE);
    refresh();
}

int GetCharacter() {
    return getch();
}

void CloseKeyboard() {
    endwin();
}
