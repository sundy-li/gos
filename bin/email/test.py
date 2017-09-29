#coding:utf-8

import os,subprocess,shlex

sendmail = "gomail  -q smtp.exmail.qq.com:25 -u boomer_kefu@truxing.com -p  NyxQGrXTpjiad18H"

def mail(email, subject, msg):
	global sendmail
	aa = sendmail + " -j " + subject + " -t " + email
	args = shlex.split(aa)

	p1 = subprocess.Popen(["echo", msg], stdout=subprocess.PIPE)
	p2 = subprocess.Popen(args,stdin=p1.stdout)
	
	p2.communicate()

if __name__ == '__main__':
    mail("libenwang@truxing.com", "主题1", "fsafsfsdfdsfdsfdsfdsf")