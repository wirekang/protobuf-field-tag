import os


def pread(cmd: str):
    f = os.popen(cmd)
    r = f.read()
    f.close()
    return r


os.chdir("example")
go_out = pread("go run ./main.go")
os.chdir("..")

rst = open("README.src.md").read()


def pcmd(cmd: str):
    prefix = "$ prototag " + cmd + " \n"
    cmd = "go run ./cmd/prototag/ " + cmd
    return prefix + pread(cmd)


os.system("cp example/example.proto .")


kv = {
    "PROTO": open("example/example.proto").read(),
    "GO": open("example/main.go").read(),
    "GO_OUT": go_out,
    "CLI1": pcmd("--help"),
    "CLI2": pcmd("-j < example.proto"),
    "CLI3": pcmd("-j -p example.proto"),
    "CLI4": pcmd("-y example.proto")
}
for k in kv:
    rst = rst.replace("##"+k+"##", kv[k])


target = open("README.md", "w")
target.write(rst)
target.close()

os.system("rm example.proto")
