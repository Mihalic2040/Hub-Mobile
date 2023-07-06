package com.Mihalic2040.hub_mobile.service;

import android.app.Notification;
import android.app.NotificationChannel;
import android.app.NotificationManager;
import android.app.PendingIntent;
import android.app.Service;
import android.content.Intent;
import android.graphics.Color;
import android.os.Build;
import android.os.IBinder;
import android.util.Log;

import androidx.annotation.Nullable;
import androidx.core.app.NotificationCompat;

import com.Mihalic2040.hub_mobile.MainActivity;
import com.Mihalic2040.hub_mobile.R;

import io.grpc.ManagedChannel;

public class node extends Service {
    private static final String TAG = "NodeService";
    private static final String CHANNEL_ID = "NodeChannel";
    public String config = "{\n" +
            "    \"Host\": \"0.0.0.0\",\n" +
            "    \"Port\": \"0\",\n" +
            "    \"Rendezvous\": \"Hub\",\n" +
            "    \"DHTServer\": true,\n" +
            "    \"ProtocolId\": \"/hub/0.0.1\",\n" +
            "    \"Bootstrap\": \"/ip4/141.145.193.111/tcp/6666/p2p/12D3KooWQd1K1k8XA9xVEzSAu7HUCodC7LJB6uW5Kw4VwkRdstPE\"\n" +
            "}";

    private Thread serviceThread;
    private boolean isServiceRunning;

    @Override
    public void onCreate() {
        super.onCreate();
        Log.i(TAG, "Service onCreate");
        Log.i(TAG, "Service onStartCommand");

        if (!isServiceRunning) {
            startServiceThread();
            isServiceRunning = true;
        }

    }

    @Override
    public int onStartCommand(Intent intent, int flags, int startId) {
        Log.i(TAG, "Service onStartCommand");

        if (!isServiceRunning) {
            startServiceThread();
            isServiceRunning = true;
        }

        return START_STICKY;
    }

    @Override
    public void onDestroy() {
        super.onDestroy();
        Log.i(TAG, "Service destroyed");

        stopServiceThread();
        isServiceRunning = false;
    }

    @Nullable
    @Override
    public IBinder onBind(Intent intent) {
        return null;
    }

    private void startServiceThread() {
        serviceThread = new Thread(() -> {
            Log.i(TAG, "Service thread started");
            deamon.Service service = new deamon.Service();
            service.config(config);
            service.start();

            // Your additional logic within the service thread

            while (!Thread.currentThread().isInterrupted()) {
                // Perform any necessary background tasks here

                try {
                    Thread.sleep(1000); // Adjust the sleep interval as needed
                } catch (InterruptedException e) {
                    Log.e(TAG, "Service thread interrupted");
                    Thread.currentThread().interrupt();
                }
            }

            Log.i(TAG, "Service thread finished");
        });

        serviceThread.start();
    }

    private void stopServiceThread() {
        if (serviceThread != null) {
            serviceThread.interrupt();
            try {
                serviceThread.join();
            } catch (InterruptedException e) {
                Log.e(TAG, "Failed to stop service thread");
            }
        }
    }


}
