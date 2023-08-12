for i in $(find -type f -name build.sh);do
  echo $i
  sh $i
done
