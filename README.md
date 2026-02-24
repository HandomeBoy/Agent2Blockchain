# 项目说明
> 本项目为农产品溯源系统。以番茄为例,将具有农作物质量检测能力的智能体与Fisco Bcos进行结合。
> <img width="1912" height="1020" alt="image" src="https://github.com/user-attachments/assets/affe2996-dd3d-41eb-9aae-418931cfee4f" />

## 项目介绍
> 1. 项目中存在三个项目,分别是Solidity2Go(go项目)、get-linker(SpringBoot项目)和src(vue项目),vue项目中文件太多,只上传了主要代码,直接将文件替换新项目中的src目录,并下载依赖。  
> 2. 项目启动了三个端口服务,分别是8080、8081和5173,确保三个端口未被占用。  
> 3. Solidity2Go项目实现了调用智能合约方法的功能。文件包含了溯源合约,部署后根据实际的角色私钥替换项目中account目录中对应的私钥文件,并修改配置文件中对应的配置即可;  
>    get-linker项目实现了调用智能体api的功能,使用Coze平台作为智能体开发平台,用户自行设计智能体,并修改配置文件中的智能体id和访问令牌。详细的过程参考get-linker文件中的README文件。
>    src项目实现了调用SpringBoot项目和go项目的接口并展示的功能,接口文档位于src文件中。

## 部署指南  
### 前提条件
> 1. 提前在虚拟机中部署Fisco Bcos区块链以及节点控制台,并将Solidity2Go文件中的sol合约部署在节点控制器上,后续的过程中需要启动虚拟机服务,并且可以在主机上访问节点控制器。  
> 2. 使用Coze搭建一个智能体并发布,获取智能体ID;申请一个访问令牌bearer。将获取到的值替换application.yml文件中对应的agent_id和personal。  
> 3. 创建一个tomato数据库，按照get-linker中README文件创建数据表的代码创建Mysql数据表。

### 部署步骤
> 1. 除vue文件外,其他两个文件直接下载即可;vue文件需要创建后自行替换src。
> 2. Solidity2Go:
>      在节点控制器中创建configs.toml文件中规定的角色,导入.pem文件,修改Solidity2Go项目中的configs.toml文件中的配置,将所有的文件名和地址改为实际的文件和地址。
>      导入ca.crt,sdk.crt,sdk.key,TomatoTraceable.abi文件。
> 3. src:
>      将接口文档和vite.config.js文件移动到根目录。
> 4. get-linker:   
>      在本机上创建application.yml文件中tomato下的三个目录,收获时期的图片存储在harvest对应的目录中,加工时期的图片存储在sampling对应的目录中,并将图片命名为生成的tomato_id。   
> 5. 三个服务全部启动后,通过下图的功能获取tomato_id。<img width="1919" height="1030" alt="image" src="https://github.com/user-attachments/assets/6d6ab203-7f57-4227-a51a-419982cefafb" />
>    将对应的图片存储在对应的路径下，改成正确的命名。之后,专家节点执行检测功能时会读取对应的文件。


