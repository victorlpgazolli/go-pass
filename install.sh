PASS=$HOME/.go-pass

# rm -rf $PASS

mkdir $PASS &>/dev/null;

rm -rf $PASS;

git clone https://github.com/victorlpgazolli/go-pass $PASS;

echo "We need sudo permissions to move the executable to /usr/sbin"

sudo cp $PASS/pass /usr/sbin/pass

echo "pass installed in /usr/sbin"

