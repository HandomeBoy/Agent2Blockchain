package com.linker.tproject.service.controller.file;

import com.linker.tproject.service.resp.SetReturn;
import com.linker.tproject.service.service.IGetTomatoImg;
import io.swagger.annotations.Api;
import io.swagger.annotations.ApiOperation;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

@RequestMapping("/agent")
@RestController
@Api(tags = "智能体服务")
@Slf4j
public class ImgController {

    @Autowired
    private SetReturn setReturn;

    @Autowired
    private IGetTomatoImg getTomatoImg;

    @PostMapping("/get")
    @ApiOperation(value = "保存图片")
    public SetReturn getImg(@RequestParam String tomatoId,@RequestParam String role) {
        try {
            String result = getTomatoImg.saveHarvestImageToFile(tomatoId,role);
            if (result != null) {
                return setReturn.isSuccess("获取成功", null);
            } else {
                return setReturn.isfail("200", "获取失败", null);
            }
        } catch (Exception e) {
            log.error("获取图片时发生异常", e);
            return setReturn.isfail("500", "系统异常", null);
        }
    }

}
