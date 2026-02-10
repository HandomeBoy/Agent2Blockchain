package com.linker.tproject.api.rpc;

import com.linker.core.base.baseclass.BaseResp;
import com.linker.tproject.api.constants.RpcConstant;
import com.linker.tproject.api.model.req.TestRpcReq;
import com.linker.tproject.api.model.resp.TestRpcResp;
import org.springframework.cloud.openfeign.FeignClient;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;

import static com.linker.tproject.api.constants.RpcConstant.*;

/**
 * @author qingwei.z
 * @Date 2022/9/5 下午2:00
 * @DESC: 测试rpc接口
 */
@FeignClient(name = FEIGN_NAME, path = FEIGN_PATH, contextId = FEIGN_NAME_POINT + "ITestRpcController")
public interface ITestRpcController {

    String basePath = RpcConstant.RPC +  "/test";

    /**
     * 测试接口
     * @param testReq
     * @return
     */
    @PostMapping(basePath + "/getTest")
    BaseResp<TestRpcResp> getTest(@RequestBody TestRpcReq testReq);
}
