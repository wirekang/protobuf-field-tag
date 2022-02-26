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
}
for k in kv:
    rst = rst.replace("##"+k+"##", kv[k])

clis = [pcmd("--help"),
        pcmd("-j < example.proto"),
        pcmd("-jp example.proto"),
        pcmd("-y example.proto"),
        pcmd("-jpa example.proto")]

cli = "".join(map(lambda c: "```bash\n"+c+"\n```\n", clis))

rst = rst.replace("##CLI##", cli)

target = open("README.md", "w")
target.write(rst)
target.close()

os.system("rm example.proto")
