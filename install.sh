# if [ "$EUID" -ne 0 ]
#   then echo "Please run as root"
#   exit
# fi

export PASS=~/.go-pass

mkdir $PASS &>/dev/null;

rm -rf go-pass;

git clone https://github.com/victorlpgazolli/go-pass;

mv go-pass/pass $PASS;

rm -rf go-pass;

echo 'pass installed!'