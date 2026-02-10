package com.linker.tproject.api.model.resp;

import com.linker.core.base.baseclass.BaseDTO;
import com.linker.core.base.baseclass.BaseResp;
import lombok.Data;

/**
 * @author qingwei.z
 * @Date 2023/4/26 下午1:00
 * @DESC:
 */
@Data
public class TestResp extends BaseDTO {

    private String name;

    private int age;

    private String school;
}
