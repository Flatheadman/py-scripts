import os
from pathlib import Path
import magika
from pprint import pprint

m = magika.Magika()




def extract_code(repo_dir, exclude_dirs=None, exclude_files=None, no_code=False):
    """Create an index, extract content of code/text files, ignoring .git directory."""
    exclude_dirs = exclude_dirs or []
    exclude_files = exclude_files or []

    code_index = []
    code_text = ""
    for root, dirs, files in os.walk(repo_dir):
        # 忽略特定目录
        for dir_to_exclude in exclude_dirs:
            if dir_to_exclude in dirs:
                dirs.remove(dir_to_exclude)

        # 忽略特定文件
        for file_to_exclude in exclude_files:
            if file_to_exclude in files:
                files.remove(file_to_exclude)

        for file in files:
            file_path = os.path.join(root, file)
            relative_path = os.path.relpath(file_path, repo_dir)
            code_index.append(relative_path)

            if no_code: # 如果不需要提取代码内容，直接跳过
                continue

            file_type = m.identify_path(Path(file_path))
            if file_type.output.group in ("text", "code"):
                try:
                    with open(file_path, "r", encoding="utf-8") as f:
                        code_text += f"----- File: {relative_path} -----\n"
                        code_text += f.read()
                        code_text += "\n-------------------------\n"
                except FileNotFoundError:
                    print(f"Error: File not found: {file_path}")
                except UnicodeDecodeError:
                    print(f"Error: Unable to decode file {file_path} with UTF-8 encoding.")
                except Exception as e:  # 捕获其他异常
                    print(f"Error reading file {file_path}: {e}")

    print("--项目结构--")
    pprint(code_index)
    print("--文件数目--")
    pprint(len(code_index))
    print("--预估token数--")
    pprint(len(code_text)/4)

    return code_index, code_text


# clone_repo(repo_url, repo_dir)

# code_index, code_text = extract_code(repo_dir)


