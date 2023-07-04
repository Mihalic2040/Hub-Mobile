package com.Mihalic2040.hub_mobile;

import androidx.appcompat.app.AppCompatActivity;

import android.os.Bundle;
import deamon.Service;

public class MainActivity extends AppCompatActivity {


    private String Config = "{\n" +
            "    \"Host\": \"0.0.0.0\",\n" +
            "    \"Port\": \"0\",\n" +
            "    \"Rendezvous\": \"Hub\",\n" +
            "    \"DHTServer\": true,\n" +
            "    \"ProtocolId\": \"/hub/0.0.1\",\n" +
            "    \"Bootstrap\": \"/ip4/141.145.193.111/tcp/6666/p2p/12D3KooWQd1K1k8XA9xVEzSAu7HUCodC7LJB6uW5Kw4VwkRdstPE\"\n" +
            "}";




    @Override
    protected void onCreate(Bundle savedInstanceState) {
        Service service = new Service();

        service.config(Config);

        service.start();

        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        while (
            service.test();
        )


        //System.out.println(service.socket());
    }
}