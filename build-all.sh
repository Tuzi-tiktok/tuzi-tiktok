for i in $(find -type f -name build.sh); do
  echo "$(basename $(dirname $i)) Building"
  sh $i
done
