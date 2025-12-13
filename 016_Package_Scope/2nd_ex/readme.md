● Package scope হচ্ছে সেই Scope যেখানে ভেরিয়েবল/ফাংশন/টাইপ পুরো প্যাকেজে ব্যবহারযোগ্য হয়। 

● Lowercase-এ শুরু হলে unexported, Uppercase-এ হলে xported (অন্য প্যাকেজ থেকে অ্যাক্সেসযোগ্য)। 



for initialize module ...

```
go mod init <module name>
```

```go mod init math_mod```
        -> it create go.mod with module name math_mod

<br>
-> for import 

```
custompackage "math_mod/CustomPackage"
custompackage.<function/ variable>
```