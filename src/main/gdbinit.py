import os

print("1.设置当前终端位gdb输出终端(and insert Pwngdb)")
print("2.Just use pwndbg")
print("3.恢复默认设置")


choice = int(input("input:"))

res = os.popen("tty")
tty = res.read()
tty_number = tty[-2]

res = os.popen("whoami")
whoami = res.read()[:-1]

gdbinit = open(f'/home/{whoami}/.gdbinit', 'w', encoding='utf-8')
cmd_str = "set context-output /dev/pts/" + str(tty_number)
if(choice == 1):
	# contain = "source ~/pwndbg/gdbinit.py\nsource ~/Pwngdb/pwngdb.py\nsource ~/Pwngdb/angelheap/gdbinit.py\n\nset disassembly-flavor intel\n" + cmd_str
	contain = f'''source ~/pwndbg/gdbinit.py
source ~/Pwngdb/pwngdb.py
source ~/Pwngdb/angelheap/gdbinit.py

define hook-run
python
import angelheap
angelheap.init_angelheap()
end
end

set disassembly-flavor intel
set context-output /dev/pts/{tty_number}
	'''

	print(contain)
	gdbinit.write(contain)
	print("set successfully!")

elif(choice == 2):
	contain = f'''source ~/pwndbg/gdbinit.py
set disassembly-flavor intel
set context-output /dev/pts/{tty_number}
	'''

	print(contain)
	gdbinit.write(contain)
	print("set successfully!")
elif(choice == 3):
	# contain = "source ~/pwndbg/gdbinit.py\nsource ~/Pwngdb/pwngdb.py\nsource ~/Pwngdb/angelheap/gdbinit.py\n\nset disassembly-flavor intel\n"
	contain = f'''source ~/pwndbg/gdbinit.py
source ~/Pwngdb/pwngdb.py
source ~/Pwngdb/angelheap/gdbinit.py

define hook-run
python
import angelheap
angelheap.init_angelheap()
end
end

set disassembly-flavor intel
	'''
	gdbinit.write(contain)
	print("recover successfully!")
else:
	print("您的输入有误！！！")