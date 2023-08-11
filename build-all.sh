for i in $(find -type f -executable -name build.sh); do
  echo $i
  sh $i
done
