# 项目说明
> 本项目为番茄检测智能体项目,通过调用api接口实现了将番茄图片传入智能体，并接收智能体返回的图片检测结果的功能。

## 模块划分
> linker-template-project-api 对外接口层，提供对外接口，model模型  
  linker-template-project-service 实际实现层，业务的具体实现

### linker-template-project-api模块
> 提供rpc的对外接口依赖包、以及对外依赖的model

### linker-template-project-service模块

#### 模块分层要求

controller -> service -> mapper -> model -> req -> resp

##### controller层
- controller层主要提供对外的接口。
- 主要包括接口定义和参数定义,部分参数禁止用户自定义,因此此处的参数仅为允许用户自定义的参数,项目默认的参数在对应的service类或配置文件中定义;
- 项目主要分为三类服务：数据持久化服务,智能体服务以及数据在智能体和数据库之前传输的中间服务(以下简称中间服务);
- 智能体服务: 包括智能体相关服务、会话相关服务、消息相关服务以及对话相关服务;
- 持久化服务: 包括数据库操作服务;
-    数据表创建代码:
-       CREATE TABLE `tomato_info` (
          `tomato_id` VARCHAR(255) NOT NULL COMMENT '番茄ID',
          `production_area` VARCHAR(255) DEFAULT NULL COMMENT '产地', 
          `harvest_date` DATE DEFAULT NULL COMMENT '收获日期',
          `processing_date` DATE DEFAULT NULL COMMENT '加工日期',
          `transport_info` VARCHAR(500) DEFAULT NULL COMMENT '运输信息',
          `harvest_image` MEDIUMBLOB COMMENT '收获图片',
          `sampling_image` MEDIUMBLOB COMMENT '采样图片',
          `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
          `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
          `harvest_evaluation` TEXT COMMENT '收获评估',
          `sampling_evaluation` TEXT COMMENT '采样评估',
          PRIMARY KEY (`tomato_id`),
          INDEX `idx_harvest_date` (`harvest_date`),
          INDEX `idx_production_area` (`production_area`),
          INDEX `idx_created_at` (`created_at`)
       ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='番茄信息表';
- 中间服务: 包括数据在智能体和数据库之间传输的转换服务,主要包括图片的转换服务:数据库存取-转换成文件-上传到智能体;
- 示例操作: 
-    1.将收获图片和加工图片放在application配置文件中声明的路径，并命名成tomato_id.jpg;
-    2.调用创建番茄信息接口,示例参数:
-      {
         "harvestDate": "2024-03-15",
         "harvestEvaluation": null,
         "harvestImage": "",
         "processingDate": "2024-03-17",
         "productionArea": "新疆塔里木盆地",
         "samplingEvaluation": null,
         "samplingImage": "",
         "tomatoId": "test01",
         "transportInfo": "冷链运输，温度保持0-4°C，运输时长24小时"
      }
-    3.调用番茄检测接口,输入身份和番茄id,返回结果为图片检测结果;


##### service层
- service层主要提供业务逻辑处理。
- 主要提供业务逻辑处理。对外部api接口的调用也是在这个层中;
- 大致逻辑为:封装请求体、封装请求头、调用api、封装返回结果;
- 本层同controller层一样,主要分为三类服务的业务逻辑处理,不再赘述;


##### mapper层
- 数据表对应的entity对象,主要用于数据库操作;


##### model层
- 其他模块中的中间对象,如参数对象、返回对象;

##### req和resp层
- req和resp层主要提供接口参数定义和返回结果定义,主要针对controller层和service层中的参数和返回结果,以适应api接口规范;
