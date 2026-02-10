package com.linker.tproject.api.model.req;

import com.linker.core.base.baseclass.BaseReq;
import com.linker.core.base.baseclass.page.BasePaginReq;
import lombok.Data;

/**
 * @author qingwei.z
 * @Date 2023/4/26 上午11:18
 * @DESC:
 */
@Data
public class TestPageReq extends BasePaginReq {

    private String name;
}
