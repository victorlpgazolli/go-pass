echo "We need root permissions to move the executable to /usr/sbin"

if [ "$EUID" -ne 0 ]
  then echo "Please run as root"
  exit
fi

PASS=$HOME/.go-pass

rm -rf $PASS

mkdir $PASS &>/dev/null;

git clone https://github.com/victorlpgazolli/go-pass $PASS;

cp $PASS/pass /usr/sbin/pass

echo "pass installed in /usr/sbin"

