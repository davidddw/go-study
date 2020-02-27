# 跨平台编译

## 编译Linux下可执行文件

```bash
yum install git gcc gcc-go rpm-build mingw64-gcc mingw64-gcc-c++
rpm -ivh oracle-instantclient12.2*.rpm
echo '/usr/lib/oracle/12.2/client64/lib' >> /etc/ld.so.conf
ldconfig
```

```bash
export PKG_CONFIG_PATH="/usr/lib64/pkgconfig"
mkdir $PKG_CONFIG_PATH
echo -e "\nName: oci8\nDescription: oci8\nLibs: -L/usr/lib/oracle/12.2/client64/lib/ -lclntsh\nCflags: -I/usr/include/oracle/12.2/client64/\nVersion: 12.2\n" > $PKG_CONFIG_PATH/oci8.pc
```

```bash
go build -v -o myprogram mypackage
```

## 编译Windows下可执行文件

```bash
mkdir -p /usr/local/oracle/
unzip instantclient_12_2.zip -d /usr/local/oracle/

sed -i 's/typedef unsigned _int64 ubig_ora;/typedef unsigned __int64 ubig_ora;/g' /usr/local/oracle/instantclient_12_2/sdk/include/oratypes.h
sed -i 's/typedef   signed _int64 sbig_ora;/typedef   signed __int64 sbig_ora;/g' /usr/local/oracle/instantclient_12_2/sdk/include/oratypes.h

export CGO_ENABLED=1
export CC=x86_64-w64-mingw32-gcc
export CXX=x86_64-w64-mingw32-g++
export CGO_CFLAGS="-I/usr/local/oracle/instantclient_12_2/sdk/include/"
export CGO_LDFLAGS="-L/usr/local/oracle/instantclient_12_2 -L/usr/local/oracle/instantclient_12_2_64/sdk/lib/msvc -lstdc++ -loci"
GOOS=windows GOARCH=amd64 go build -v

# 编译linux
export CGO_ENABLED=
export CC=
export CXX=
export CGO_CFLAGS=
export CGO_LDFLAGS=
go build 
```