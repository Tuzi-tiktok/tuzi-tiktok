for i in $(find -type f -name build.sh);do
  echo $(basename $i)
  sh $i
done
