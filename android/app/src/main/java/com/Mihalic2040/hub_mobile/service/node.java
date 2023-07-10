package com.Mihalic2040.hub_mobile.service;

import android.app.Notification;
import android.app.NotificationChannel;
import android.app.NotificationManager;
import android.app.PendingIntent;
import android.app.Service;
import android.content.Context;
import android.content.Intent;
import android.graphics.Color;
import android.net.wifi.WifiManager;
import android.os.Build;
import android.os.IBinder;
import android.util.Log;

import androidx.annotation.Nullable;
import androidx.core.app.NotificationCompat;

import com.Mihalic2040.hub_mobile.MainActivity;
import com.Mihalic2040.hub_mobile.R;


public class node extends Service {
    private static final String TAG = "NodeService";
    private static final String CHANNEL_ID = "NodeChannel";
    private static final int NOTIFICATION_ID = 1;
    private static final String NOTIFICATION_TITLE = "Node Service";
    private static final String NOTIFICATION_TEXT = "Running";

    private String config = "{\n" +
            "    \"Host\": \"0.0.0.0\",\n" +
            "    \"Port\": \"0\",\n" +
            "    \"Secret\": \"12121\",\n" +
            "    \"Rendezvous\": \"Hub\",\n" +
            "    \"ProtocolId\": \"/hub/0.0.1\",\n" +
            "    \"Bootstrap\": \"/ip4/141.145.193.111/tcp/6666/p2p/12D3KooWQd1K1k8XA9xVEzSAu7HUCodC7LJB6uW5Kw4VwkRdstPE\"\n" +
            "}";

    private Thread serviceThread;
    private boolean isServiceRunning;

    @Override
    public void onCreate() {
        super.onCreate();
        Log.i(TAG, "Service onCreate");
        createNotificationChannel();
        startServiceThread();
    }

    @Override
    public int onStartCommand(Intent intent, int flags, int startId) {
        Log.i(TAG, "Service onStartCommand");

        if (!isServiceRunning) {
            startServiceThread();
        }

        // Return START_REDELIVER_INTENT to ensure the service is restarted with the last intent
        return START_REDELIVER_INTENT;
    }

    @Override
    public void onDestroy() {
        super.onDestroy();
        Log.i(TAG, "Service onDestroy");

        stopServiceThread();
    }

    @Nullable
    @Override
    public IBinder onBind(Intent intent) {
        return null;
    }

    private void createNotificationChannel() {
        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.O) {
            NotificationChannel channel = new NotificationChannel(
                    CHANNEL_ID,
                    "Node Channel",
                    NotificationManager.IMPORTANCE_DEFAULT
            );
            channel.enableLights(false);
            channel.enableVibration(false);

            NotificationManager notificationManager = getSystemService(NotificationManager.class);
            notificationManager.createNotificationChannel(channel);
        }
    }

    public void acquireMulticastLock() {
        WifiManager wifiManager = (WifiManager) this.getSystemService(Context.WIFI_SERVICE);
        WifiManager.MulticastLock multicastLock = wifiManager.createMulticastLock("myMulticastLockTag");
        multicastLock.acquire();
    }

    public void releaseMulticastLock() {
        WifiManager wifiManager = (WifiManager) this.getSystemService(Context.WIFI_SERVICE);
        WifiManager.MulticastLock multicastLock = wifiManager.createMulticastLock("myMulticastLockTag");
        multicastLock.release();
    }

    private void startServiceThread() {
        serviceThread = new Thread(() -> {
            Log.i(TAG, "Service thread started");

            // Additional logic within the service thread
            acquireMulticastLock();
            deamon.Service service = new deamon.Service();
            service.config(config);
            service.start();



            while (!Thread.currentThread().isInterrupted()) {
                // Perform any necessary background tasks here

                try {
                    Thread.sleep(1000); // Adjust the sleep interval as needed
                    service.requestTest();
                } catch (InterruptedException e) {
                    Log.e(TAG, "Service thread interrupted");
                    releaseMulticastLock();
                    Thread.currentThread().interrupt();
                }
            }

            Log.i(TAG, "Service thread finished");
        });

        serviceThread.start();

        // Show a foreground notification
        Notification notification = buildNotification();
        startForeground(NOTIFICATION_ID, notification);
        isServiceRunning = true;
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

        // Stop the service as foreground and remove the notification
        stopForeground(true);
        isServiceRunning = false;
    }

    private Notification buildNotification() {
        Intent notificationIntent = new Intent(this, MainActivity.class);
        PendingIntent pendingIntent = PendingIntent.getActivity(this, 0, notificationIntent, PendingIntent.FLAG_IMMUTABLE);

        NotificationCompat.Builder builder = new NotificationCompat.Builder(this, CHANNEL_ID)
                .setContentTitle(NOTIFICATION_TITLE)
                .setContentText(NOTIFICATION_TEXT)
                .setSmallIcon(R.drawable.ic_launcher_foreground)
                .setContentIntent(pendingIntent)
                .setOngoing(true);

        return builder.build();
    }
}
