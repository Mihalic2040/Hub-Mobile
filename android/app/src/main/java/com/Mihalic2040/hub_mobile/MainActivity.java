package com.Mihalic2040.hub_mobile;

import androidx.appcompat.app.AppCompatActivity;

import android.content.Context;
import android.content.Intent;
import android.content.SharedPreferences;
import android.os.Bundle;
import android.util.Log;
import android.widget.Button;

import deamon.Service;
import com.Mihalic2040.hub_mobile.service.node;

import com.Bridge.api.Peers;

public class MainActivity extends AppCompatActivity {

    private String Config = "{\n" +
            "    \"Host\": \"0.0.0.0\",\n" +
            "    \"Port\": \"0\",\n" +
            "    \"Rendezvous\": \"Hub\",\n" +
            "    \"DHTServer\": true,\n" +
            "    \"ProtocolId\": \"/hub/0.0.1\",\n" +
            "    \"Bootstrap\": \"/ip4/141.145.193.111/tcp/6666/p2p/12D3KooWQd1K1k8XA9xVEzSAu7HUCodC7LJB6uW5Kw4VwkRdstPE\"\n" +
            "}";


    private Intent serviceIntent;
    private boolean isServiceRunning;


    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        Button startButton = findViewById(R.id.start_btn);
        startButton.setOnClickListener(v -> startNodeService());


        Button stopButton = findViewById(R.id.stop_btn);
        stopButton.setOnClickListener(v -> stopNodeService());

        serviceIntent = new Intent(this, node.class);
        if (isServiceRunning == false) {
            startService(serviceIntent);
            isServiceRunning = true;
        }

        // Restore the service state from preferences


        new Thread(() -> {
            while (true) {
                try {
                    // Sleep for 20 seconds
                    Thread.sleep(2000);

                    // Create an instance of Peers class
                    Peers requestHandler = new Peers();

                    // Call the request() method
                    requestHandler.request();
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            }

        }).start();



        //System.out.println(service.socket());
    }


    private void startNodeService() {
        if (!isServiceRunning) {
            startService(serviceIntent);
            isServiceRunning = true;
            Log.i("MainActivity", "Node Service started");
        }
    }

    private void stopNodeService() {
        if (isServiceRunning) {
            stopService(serviceIntent);
            isServiceRunning = false;
            Log.i("MainActivity", "Node Service stopped");
        }
    }

}