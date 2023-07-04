package com.Mihalic2040.hub_mobile;

import androidx.appcompat.app.AppCompatActivity;

import android.os.Bundle;
import hub.App;

public class MainActivity extends AppCompatActivity {

    private App app;

    private String Config = "{\n" +
            "    \"Host\": \"0.0.0.0\",\n" +
            "    \"Port\": \"0\",\n" +
            "    \"Rendezvous\": \"Hub\",\n" +
            "    \"DHTServer\": true,\n" +
            "    \"ProtocolId\": \"/hub/0.0.1\",\n" +
            "    \"Bootstrap\": \"/ip4/141.145.193.111/tcp/6666/p2p/12D3KooWQd1K1k8XA9xVEzSAu7HUCodC7LJB6uW5Kw4VwkRdstPE\"\n" +
            "}";

    private String request = "{\n" +
            "  \"User\": \"12D3KooWGQ4ncdUVMSaVrWrCU1fyM8ZdcVvuWa7MdwqkUu4SSDo4\",\n" +
            "  \"Payload\": \"hhh\",\n" +
            "  \"Handler\": \"MyHandler\"\n" +
            "}";


    @Override
    protected void onCreate(Bundle savedInstanceState) {
        app = new App();



        try {
            app.configFromJson(Config);
        } catch (Exception e) {
            // Handle the exception here
            e.printStackTrace();
        }
        app.start(true);

        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        try {
            Thread.sleep(3000);
            app.newRequest(request);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}