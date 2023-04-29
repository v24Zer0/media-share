package com.mediashare.authenticationservice;

import com.mediashare.authenticationservice.util.UUIDRegex;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import static org.assertj.core.api.Assertions.assertThat;

public class UUIDRegexTests {

    private UUIDRegex uuidRegex;

    @BeforeEach
    public void setup() {
        this.uuidRegex = new UUIDRegex();
    }

    @Test
    @DisplayName("Should correctly match a valid uuid")
    public void testMatchID() {
        String id = "c73bcdcc-2669-4bf6-81d3-e4ae73fb11fd";
        boolean ok = this.uuidRegex.validateID(id);

        assertThat(ok).isTrue();
    }

    @Test
    @DisplayName("Should not match an invalid uuid")
    public void testMatchIDWithInvalidUUID() {
        String id1 = "c73bcdcc26694bf681d3e4ae73fb11fd";
        boolean ok = this.uuidRegex.validateID(id1);
        assertThat(ok).isFalse();


        String id2 = "definitely-not-a-uuid";
        ok = this.uuidRegex.validateID(id2);
        assertThat(ok).isFalse();
    }
}
