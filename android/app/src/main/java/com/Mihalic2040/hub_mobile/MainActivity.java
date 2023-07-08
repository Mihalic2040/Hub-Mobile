package com.Mihalic2040.hub_mobile;

import androidx.appcompat.app.AppCompatActivity;

import android.content.Intent;
import android.os.Bundle;
import android.provider.Settings;
import android.widget.Button;
import android.widget.TextView;

import com.Bridge.protocols.PeersResponse;
import com.Mihalic2040.hub_mobile.service.node;

import com.Bridge.api.Peers;

public class MainActivity extends AppCompatActivity {
    private Intent serviceIntent;

    private TextView peersTextView;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        //init SERVICE
        serviceIntent = new Intent(this, node.class);

        Button startButton = findViewById(R.id.start_btn);
        startButton.setOnClickListener(v -> startNodeService());


        Button stopButton = findViewById(R.id.stop_btn);
        stopButton.setOnClickListener(v -> stopNodeService());

        peersTextView = findViewById(R.id.peers);

        // Restore the service state from preferences





    }


    private void testNodeConn() {
        new Thread(() -> {
            while (true) {
                try {
                    // Sleep for 20 seconds
                    Thread.sleep(2000);

                    // Create an instance of Peers class
                    Peers requestHandler = new Peers();

                    // Call the request() method
                    PeersResponse res = requestHandler.request();

                    runOnUiThread(() -> peersTextView.setText(res.getPeers()));
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            }

        }).start();
    }

    private void startNodeService() {
        // Start the service
        startService(serviceIntent);

        try {
            Thread.sleep(5000);
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
        testNodeConn();
    }

    private void stopNodeService() {
        // Stop the service
        stopService(serviceIntent);
    }

}