#!/bin/bash

# 输入文件路径
input_file="./README.md"

# 输出文件路径
output_file="./output.md"

# 逐行读取输入文件，并在每行的末尾添加两个空格后写入输出文件
while IFS= read -r line
do
  echo "$line  " >> "$output_file"
done < "$input_file"

echo "处理完成！输出文件为 $output_file"