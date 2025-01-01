# -*- coding: utf-8 -*-
import repository

#################################################
title = "dss-src" # 项目标题
repo_dir = "/root/MyProjects/solidity/dss/src"  
no_code = False # 是否提取代码内容，如果为 True 则只提取代码文件索引

exclude_dirs_project = ['deployments','deployments-zk','TetherTokenV2.sol']
exclude_files_project = ['CHANGELOG.md',]

exclude_dirs= exclude_dirs_project + [".git",".github", "artifacts", "cache", "node_modules", "typechain-types",'.devcontainer','.yarn','.changeset','.husky','audits'] 
exclude_files= exclude_files_project + [".gitignore", ".dockerignore", "LICENSE", "package-lock.json", "yarn.lock"]
#################################################


code_index, code_text = repository.extract_code(repo_dir, exclude_dirs, exclude_files, no_code)

if no_code:
    print("代码内容已经忽略，仅提取了代码文件索引。")
    # 将 code_index 以覆盖的方式写入 output.txt 文件，指定 UTF-8 编码
    with open("outputs/"+title+"_index.txt", "w", encoding="utf-8") as f:
        f.write("\n".join(code_index))
    print("done")
else:
    # 模型系统指令
    system_instruction = f"""
    你是一个资深代码专家，在接下来的一系列问题中，你将帮助我分析这个代码库，深入地学习它的设计与实现。
    所以请你记住这个代码库，因为你将在接下来的问题中回答关于这个代码库的问题。

    整个代码库的信息如下:
    - 整个代码库的文件索引如下：
        \n\n{code_index}\n\n.
    - 所有项目文件的内容已经拼接在一起，所有的项目代码如下：
        \n\n{code_text}\n\n

    """

    # 将 system_instruction 以覆盖的方式写入 output.txt 文件，指定 UTF-8 编码
    with open("outputs/"+title+".txt", "w", encoding="utf-8") as f:
        f.write(system_instruction)

    print("done")