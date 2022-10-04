#include<stdio.h>
#include<stdlib.h>
#include<string.h>
int main(){
// 	FILE* fp;
// 	char buffer[100];
// 	fp = popen("pwd", "r");
// 	fgets(buffer, sizeof(buffer), fp);
// 	char cmd[100] = "python3 ";
// 	strcat(cmd, buffer);
	
// 	cmd[strlen(cmd) - 1] = 0;

// 	strcat(cmd, "/gdbinit.py");
// 	printf("%s", cmd);
	system("python3 /home/kali/Desktop/git/gdbinit/gdbinit.py");
	return 0;
}
