package com.linker.tproject.service.controller.file;

import com.linker.tproject.service.entity.TomatoInfo;
import com.linker.tproject.service.resp.SetReturn;
import com.linker.tproject.service.service.impl.database.CreateTomatoService;
import io.swagger.annotations.Api;
import io.swagger.annotations.ApiOperation;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

@Api(tags = "番茄信息")
@RestController
@RequestMapping("/tomato")
@Slf4j
public class CreateTomatoController {

    /**
     * 修复后的 createTomato 方法，正确处理 JSON 数据
     */
    @Autowired
    CreateTomatoService createTomatoService;

    @PostMapping("/create")
    @ApiOperation(value = "创建番茄信息")
    public SetReturn createTomato(@RequestBody TomatoInfo tomatoInfo) {
        log.info("创建番茄信息");

        try {
            // 验证必填字段
            if (tomatoInfo.getTomatoId() == null || tomatoInfo.getTomatoId().isEmpty()) {
                throw new IllegalArgumentException("tomato_id 不能为空");
            }

            // 调用服务层方法创建数据
            Boolean result = createTomatoService.createTomato(
                    tomatoInfo.getTomatoId(),
                    tomatoInfo.getProductionArea(),
                    tomatoInfo.getHarvestDate(),
                    tomatoInfo.getProcessingDate(),
                    tomatoInfo.getTransportInfo(),
                    tomatoInfo.getCreatedAt()
            );

            return new SetReturn().isSuccess("创建信息成功",result);
        } catch (IllegalArgumentException e) {
            return new SetReturn().isfail("200","参数错误",e.getMessage());
        }
    }


}
