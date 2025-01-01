import sys
import random
import config.random_learn as config


def print_random_dict(my_dict):
  print(my_dict['key'])
  for index, value in enumerate(my_dict['value']):
    print(f"{index}: {value}")
  

def print_random_element(array_name=None, arrays=config.arrays):

  if array_name is None:
    # 没有参数，从所有数组变量的元素中随机打印一个
    all_elements = []
    for name, array in arrays.items():
      all_elements.extend([(name, element) for element in array])  # 将数组名和元素一起存储
    chosen_array_name, element = random.choice(all_elements)  # 随机选择一个元素和对应的数组名
    print(chosen_array_name)
    print_random_dict(element)
  else:
    # 检查数组变量是否存在
    if array_name not in arrays:
      print(f"Error: 数组变量 '{array_name}' 不存在")
      return

    # 获取数组变量
    array = arrays[array_name]

    # 打印随机元素和数组名
    print(array_name)
    print_random_dict(random.choice(array))

if __name__ == "__main__":
  # 获取命令行参数
  if len(sys.argv) > 2:
    print("Usage: python main.py [<array_name>]")
    sys.exit(1)

  if len(sys.argv) == 2:
    array_name = sys.argv[1]
  else:
    array_name = None

  print_random_element(array_name)