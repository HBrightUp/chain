#环境
Python version 3.68 及其以上

#构建
pip3  install -r requirements.txt

#运行
python3 ./src/main.py

#打包
pyinstaller -F src/main.py

#注意事项
1，所有的空投使用csv格式文件
2，所有的配置使用json文件
3，批量文件使用xls格式
