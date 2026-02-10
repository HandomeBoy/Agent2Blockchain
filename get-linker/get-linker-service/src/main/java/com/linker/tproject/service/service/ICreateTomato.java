package com.linker.tproject.service.service;

import com.linker.tproject.service.entity.TomatoInfo;

public interface ICreateTomato {

    Boolean createTomato(String tomato_id, String production_area,
                            java.sql.Date harvest_date, java.sql.Date processing_date, String transport_info,
                            java.sql.Timestamp created_at);

}
